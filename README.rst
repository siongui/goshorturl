================
Golang Short URL
================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/goshorturl?status.svg
   :target: https://godoc.org/github.com/siongui/goshorturl

.. image:: https://github.com/siongui/goshorturl/workflows/ci/badge.svg
    :target: https://github.com/siongui/goshorturl/blob/master/.github/workflows/ci.yml

.. image:: https://goreportcard.com/badge/github.com/siongui/goshorturl
   :target: https://goreportcard.com/report/github.com/siongui/goshorturl

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://github.com/siongui/goshorturl/blob/master/UNLICENSE


To build the short link engine by Go_.

Development Environment:

  - `Ubuntu 20.04`_
  - `Go 1.17.1`_


Requirement
+++++++++++

- Console app, receive “URL” from any website.
- Implement logic to make a unique shortlink and store in database. Maps the
  original link with short link generated by system.
- The application is able to translate the short link to be original link. This
  can be separate 2 menu / inputs.
- Use Postgres DB, bun_ to connect DB.
- Unit Testing is required.


Solution Concept
++++++++++++++++

The range of Go *uint64* is `[0, 18 446 744 073 709 551 615]`_, so we assume
that at most 18,446,744,073,709,551,615 short links can be stored.

Each record in the table of database contains three value
**(id uint64, originalURL string, shortURL string)**.

- For every **originalURL** is received, search **originalURL** field in the
  table of database:
  * If record found, return **shortURL**.
  * If record not found, randomly generate a *uint64* value. Search the value
    in the table of database. If the value already exists, re-genetate a
    *uint64* value until the value does not exist in the table of database. Then
    use base58_ to encode the value into a string. This string is the
    **shortURL** we need.

- For every **shortURL** is received, search **shortURL** field in the table of
  database:
  * If record found, return **originalURL**.
  * If record not found, return HTTP 404.


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] | `algorithm for url shortening - Google search <https://www.google.com/search?q=algorithm+for+url+shortening>`_
       | `algorithm for url shortening - DuckDuckGo search <https://duckduckgo.com/?q=algorithm+for+url+shortening>`_
       | `algorithm for url shortening - Ecosia search <https://www.ecosia.org/search?q=algorithm+for+url+shortening>`_
       | `algorithm for url shortening - Qwant search <https://www.qwant.com/?q=algorithm+for+url+shortening>`_
       | `algorithm for url shortening - Bing search <https://www.bing.com/search?q=algorithm+for+url+shortening>`_
       | `algorithm for url shortening - Yahoo search <https://search.yahoo.com/search?p=algorithm+for+url+shortening>`_
       | `algorithm for url shortening - Baidu search <https://www.baidu.com/s?wd=algorithm+for+url+shortening>`_
       | `algorithm for url shortening - Yandex search <https://www.yandex.com/search/?text=algorithm+for+url+shortening>`_

.. [2] `How to design a tiny URL or URL shortener? - GeeksforGeeks <https://www.geeksforgeeks.org/how-to-design-a-tiny-url-or-url-shortener/>`_

.. [3] | `golang url shortener - Google search <https://www.google.com/search?q=golang+url+shortener>`_
       | `golang url shortener - DuckDuckGo search <https://duckduckgo.com/?q=golang+url+shortener>`_
       | `golang url shortener - Ecosia search <https://www.ecosia.org/search?q=golang+url+shortener>`_
       | `golang url shortener - Qwant search <https://www.qwant.com/?q=golang+url+shortener>`_
       | `golang url shortener - Bing search <https://www.bing.com/search?q=golang+url+shortener>`_
       | `golang url shortener - Yahoo search <https://search.yahoo.com/search?p=golang+url+shortener>`_
       | `golang url shortener - Baidu search <https://www.baidu.com/s?wd=golang+url+shortener>`_
       | `golang url shortener - Yandex search <https://www.yandex.com/search/?text=golang+url+shortener>`_

.. [4] `How to Create a Custom URL Shortener Using Golang and Redis <https://intersog.com/blog/how-to-write-a-custom-url-shortener-using-golang-and-redis/>`_

.. [5] | `Let's build a URL shortener in Go - Final Part : Forwarding <https://www.eddywm.com/lets-build-a-url-shortener-in-go-part-iv-forwarding/>`_
       | `GitHub - eddywm/go-shortener-wm: A  super fast url shortener service written in Go <https://github.com/eddywm/go-shortener-wm>`_

.. _Go: https://golang.org/
.. _Ubuntu 20.04: https://releases.ubuntu.com/20.04/
.. _Go 1.17.1: https://golang.org/dl/
.. _UNLICENSE: https://unlicense.org/
.. _bun: https://github.com/uptrace/bun
.. _[0, 18 446 744 073 709 551 615]: https://stackoverflow.com/a/6878625
.. _base58: https://github.com/itchyny/base58-go