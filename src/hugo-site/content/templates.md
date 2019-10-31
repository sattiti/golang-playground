---
title: "Templates"
h1: "Templates"
date: 
draft: false
linkTitle: "Templates"
weight: 1
---

## Template lookup order

### Single Page(Leaf page)

#### Single page in section
- /layouts
    1. /posts
        1. /single.html.html
        2. /single.html
    2. /_default
        1. /single.html.html
        2. /single.html

#### Single page in section with layout set
- /layouts
    1. /posts
        1. /demolayout.html.html
        2. /single.html.html
        3. /demolayout.html
        4. /single.html
    2. /_default
        1. /demolayout.html.html
        2. /single.html.html
        3. /demolayout.html
        4. /single.html

#### AMP single page
- /layouts
    1. /posts
        1. /single.amp.html
        2. /single.html
    2. /_default
        1. /single.amp.html
        2. /single.htm

#### AMP single page with lang
- /layouts
    1. /posts
        1. /single.fr.amp.html
        2. /single.amp.html
        3. /single.fr.html
        4. /single.html
    2. /_default
        1. /single.fr.amp.html
        2. /single.amp.html
        3. /single.fr.html
        4. /single.html


### Home Page(index)
#### Home page
- /layouts
    1. /
        1. index.html.html
        2. home.html.html
        3. list.html.html
        4. index.html
        5. home.html
        6. list.html
    2. /_default
        1. /index.html.html
        2. /home.html.html
        3. /list.html.html
        4. /index.html
        5. /home.html
        6. /list.html

#### Home page with type set
- /layouts
    1. /demotype
        1. /index.html.html
        2. /home.html.html
        3. /list.html.html
        4. /index.html
        5. /home.html
        6. /list.html
    2. /
        1. index.html.html
        2. home.html.html
        3. list.html.html
        4. index.html
        5. home.html
        6. list.html
    3. /_default
        1. /index.html.html
        2. /home.html.html
        3. /list.html.html
        4. /index.html
        5. /home.html
        6. /list.html

#### Home page with layout set
- layouts
    1. /
        1. demolayout.html.html
        2. index.html.html
        3. home.html.html
        4. list.html.html
        5. demolayout.html
        6. index.html
        7. home.html
        8. list.html
    2. /_default
        1. /demolayout.html.html
        2. /index.html.html
        3. /home.html.html
        4. /list.html.html
        5. /demolayout.html
        6. /index.html
        7. /home.html
        8. /list.html

#### AMP home with language
- layouts
    1. /
        1. index.fr.amp.html
        2. home.fr.amp.html
        3. list.fr.amp.html
        4. index.amp.html
        5. home.amp.html
        6. list.amp.html
        7. index.fr.html
        8. home.fr.html
        9. list.fr.html
        10. index.html
        11. home.html
        12. list.html
    2. /_default
        1. /index.fr.amp.html
        2. /home.fr.amp.html
        3. /list.fr.amp.html
        4. /index.amp.html
        5. /home.amp.html
        6. /list.amp.html
        7. /index.fr.html
        8. /home.fr.html
        9. /list.fr.html
        10. /index.html
        11. /home.html
        12. /list.html

### Section (categories)
#### Section list for section
- layouts
    1. /posts
        1. /posts.html.html
        2. /section.html.html
        3. /list.html.html
        4. /posts.html
        5. /section.html
        6. /list.html
    2. /section
        1. /posts.html.html
        2. /section.html.html
        3. /list.html.html
        4. /posts.html
        5. /section.html
        6. /list.html
    3. /_default
        1. /posts.html.html
        2. /section.html.html
        3. /list.html.html
        4. /posts.html
        5. /section.html
        6. /list.html

#### Section list for posts section with type set to blog
- /layouts
    1. /blog
        1. /posts.html.html
        2. /section.html.html
        3. /list.html.html
        4. /posts.html
        5. /section.html
        6. /list.html
    2. /posts
        1. /posts.html.html
        2. /section.html.html
        3. /list.html.html
        4. /posts.html
        5. /section.html
        6. /list.html
    3. /section
        1. /posts.html.html
        2. /section.html.html
        3. /list.html.html
        4. /posts.html
        5. /section.html
        6. /list.html
    4. /_default
        1. /posts.html.html
        2. /section.html.html
        3. /list.html.html
        4. /posts.html
        5. /section.html
        6. /list.html
