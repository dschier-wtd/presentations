---
marp: true
theme: "rose-pine"

author: "Daniel Schier"
title: "Ansible - The hard things"

paginate: true
header: "Ansible Community Day"
footer: "Ansible - The hard things"
---

## Ansible - The hard things

After over 10 years of Ansible, I might have faced some things.

---

## :calendar: Agenda

- :wave: Hi
- ⚠️ Disclaimer
- 🔄 Update handling
- 🔔 Multiple notify
- 🌐 Global become
- 📏 Side effects
- ♾️ Nested loops
- 🔒 Vault vars
- :heart: Thank you

---

## :wave: Hi

- My Name is Daniel.
- Some might say, I am a "DevOps engineer", ...
- ... but I see myself as an IT craftsman.
- I am also contributing to Open Source, democratizing software.
- When it comes to Ansible it's ~10 years of professional and community
  work.
- I like cats, cookies, computer thingies (technical term), ...
- ... and Penguins, of course.

---

## ⚠️ Disclaimer

When doing this presentation, I wasn't aware of AGP (Ansible Good Practices).
Therefore, this might or might not include things that you already know from
there.

In any case, I see some opportunity for contributions.

---

## 🔄 Update handling

Let's start with something simple - Handling updates.
Yes, this is something we might do - even I encourage to use "present" in
regular playbooks.

```yaml

- name: "Update mariadb"
  ansible.builtin.package:
    name: "mariadb-server"
    state: "latest"
  become: true

```

What can possibly go wrong?

---

## 🔔 Multiple notify 1/2

Ever wanted to trigger multiple handlers, but not call each of them? You might
have a task like:

```yaml

- name: "Manage PHP Packages"
  ansible.builtin.package:
    name: "php"
    state: "latest"
  notify:
    - "Restart LAMP"

```

Since I am lazy, I don't want to mention each and every affected handler.

---

## 🔔 Multiple notify 2/2

Now, you need multiple handlers listening to this.

```yaml
- name: "Restart httpd"
  ansible.builtin.service:
    name: "httpd.service"
    state: "restarted"
  listen: "Restart LAMP"

- name: "Restart mariadb"
  ansible.builtin.service:
    name: "mariadb.service"
    state: "restarted"
  listen: "Restart LAMP"
```

Well, this was easy.

---

## 🌐 Global become

Now, we need to become a bit controversial.

```yaml
- hosts: all
  become: true
```

or

```yaml
  tasks:
    - ansible.builtin.package:
        name: "foo"
        state: "present"
      become: true
```

And why do you think so?

---

## 📏 Side Effects

I have seen:

```yaml
- hosts: "all"

  handlers:
    - name: "Start & Enable Service"
      ansible.builtin.service:
        name: "myservice"
        state: "started"
        enabled: true
      become: true

  ...
```

How might this break?

---

## ♾️ Nested loops 1/3

Let's assume an inventory like:

```yaml
users:
  - username: "alice"
    password: "$6$password_hash"
    pub_keys:
      - "$6$one"
      - "$6$two"
  - username: "bob"
    password: "$6$another_hash"
    pub_keys:
      - "$6$three"
```

And we want to create the users. Which requires two tasks. One for
`ansible.builtin.user` and one for `ansible.builtin.authorized_key`.

---

## ♾️ Nested loops 2/3

You might consider:

> "Let's loop over the first and than the second part."

Or:

> "Let's use a block."

But that's not entirely what we want.

---

## ♾️ Nested loops 3/3

What about:

```shell
$ tree
.
├── main.yml
└── user.yml
```

And a `main.yml` like:

```yaml
- ansible.builtin.import_tasks:
    file: "user.yml"
  loop: "{{ users }}"
  loop_control:
    loop_var: "user"
    label: "{{ user.username }}"
```

This way, it's easy to add new logic and loop per user.

---

## 🔒 Vault Vars 1/3

You might find an inventory like this:

```shell
$ tree
.
└── all
    ├── vars.yml
    └── vault.yml
```

---

## 🔒 Vault Vars 2/3

And a `vars.yml` like:

```yaml
username: "bar"
```

Maybe your `vault.yml`` looks like this (decrypted):

```yaml
password: "SuperSecure"
```

Any problems here?

---

## 🔒 Vault Vars 3/3

What about a `vars.yml` like:

```yaml
username: "bar"
password: "{{ vault_password }}"
```

and a `vault.yml` like:

```yaml
vault_password: "SuperSecure"
```

Now, you don't need to decrypt the vault to understand whats going on.

---

## #️⃣ Shebang #!

Lastly, do you know that Ansible playbooks can be executed?

```yaml
#!/usr/bin/env ansible-playbook

- name: "Execute me directly"
  hosts: "all"

  vars:
    ...

  tasks:
    ...

```

Just save it in `/usr/local/bin/` and `chmod u+x`.

---

## :heart: Thank you

Thank you for your attendance, patience and input. I would love to talk with you
in the break, get your feedback and suggestions.

---

## 👤 Contact

You can find me all over the interwebs:

- 🐘 <https://fosstodon.org/@dschier>
- 🔧 <https://github.com/dschier-wtd>
- 🔧 <https://github.com/whiletruedoio>
- 🖋️ <https://blog.while-true-do.io>
- ✉️ <dschier@while-true-do.io>

Slides and example code at: <https://github.com/dschier-wtd/presentation/>.
