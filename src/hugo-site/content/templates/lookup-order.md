---
title: "Lookup Order | Templates"
h1: "Lookup Order"
date: 
draft: false
linkTitle: "Lookup Order"
weight: 1
---

## Template lookup order

### Base Template
```plain
1.
/layouts
  /section
    /<TYPE>-baseof.html

2.
/themes
  /<THEME>
    /layouts
      /section
        /<TYPE>-baseof.html

3.
/layouts
  /<TYPE>
    /baseof.html

4.
/themes
  /<THEME>
    /layouts
      /<TYPE>
        /baseof.html

5.
/layouts
  /section
    /baseof.html

6.
/themes
  /<THEME>
    /layouts
      /section
        /baseof.html

7.
/layouts
  /_default
    /<TYPE>-baseof.html

8.
/themes
  /<THEME>
    /layouts
      /_default
        /<TYPE>-baseof.html

9.
/layouts
  /_default
    /baseof.html

10.
/themes
  /<THEME>
    /layouts
      /_default
        /baseof.html
```

### Single Page(Leaf page)
#### Single page in a section
**Folder Order**
```plain
/layouts
    /posts
    /_default
```

**File Order**
```plain
single.html.html
single.html
```

#### Single page in a section with layout set
**Folder Order**
```plain
/layouts
    /posts
    /_default
```

**File Order**
```plain
demolayout.html.html
single.html.html
demolayout.html
single.html
```

#### AMP in a section single page
**Folder Order**
```plain
/layouts
    /posts
    /_default
```

**File Order**
```plain
single.amp.html
single.html
```

#### AMP in a section single page with lang
**Folder Order**
```plain
/layouts
    /posts
    /_default
```

**File Order**
```plain
single.fr.amp.html
single.amp.html
single.fr.html
single.html
```

### Home Page(index)
#### Home page
**Folder Order**
```plain
/layouts
    /
    /_default
```
**File Order**
```plain
index.html.html
home.html.html
list.html.html
index.html
home.html
list.html
```

#### Home page with type set
**Folder Order**
```plain
/layout
    /demotype
    /
    /_default
```

**File Order**
```plain
index.html.html
home.html.html
list.html.html
index.html
home.html
list.html
```

#### Home page with layout set
**Folder Order**
```plain
/layouts
    /
    /_default
```

**File Order**
```plain
demolayout.html.html
index.html.html
home.html.html
list.html.html
demolayout.html
index.html
home.html
list.html
```

#### AMP home with language
**Folder Order**
```plain
/layouts
    /
    /_default
```

**File Order**
```plain
index.fr.amp.html
home.fr.amp.html
list.fr.amp.html
index.amp.html
home.amp.html
list.amp.html
index.fr.html
home.fr.html
list.fr.html
index.html
home.html
list.html
```

### Section (categories)
#### Section list for a section
**Folder Order**
```plain
/layouts
    /posts
    /section
    /_default
```

**File Order**
```plain
posts.html.html
section.html.html
list.html.html
posts.html
section.html
list.html
```

#### Section list for posts section with type set to blog
**Folder Order**
```plain
/layouts
    /blog
    /posts
    /section
    /_default
```

**File Order**
```plain
posts.html.html
section.html.html
list.html.html
posts.html
section.html
list.html
```

### Taxonomy List Pages
#### Taxonomy list in categories
**Folder Order**
```plain
/layouts
    /categories
    /taxonomy
    /category
    /_default
```

**File Order**
```plain
category.html.html
taxonomy.html.html
list.html.html
category.html
taxonomy.html
list.html
```

#### Taxonomy term in categories
**Folder Order**
```plain
/layouts
    /categories
    /taxonomy
    /category
    /_default
```

**File Order**
```plain
category.terms.html.html
terms.html.html
list.html.html
category.terms.html
terms.html
list.html
```
