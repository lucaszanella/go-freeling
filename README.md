# go-freeling

**Natural Language Processing** in GO

This is a partial port of Freeling 3.1 (http://nlp.lsi.upc.edu/freeling/).

This is a continuation of the work from https://github.com/advancedlogic/go-freeling, where I'm adding support for all languages.

License is GPL to respect the License model of Freeling.

This is the list of features already implemented:

* Text tokenization
* Sentence splitting
* Morphological analysis
* Suffix treatment, retokenization of clitic pronouns
* Flexible multiword recognition
* Contraction splitting
* Probabilistic prediction of unknown word categories
* Named entity detection
* PoS tagging
* Chart-based shallow parsing
* Named entity classification (With an external library MITIE - https://github.com/mit-nlp/MITIE)
* Rule-based dependency parsing

-

**How to use it**:

The port where I got this from was using a web based processing example, I'm gonna write a simple parser for a sentence soon.

TODO:
* clean code
* add comments
* add tests
* implement WordNet-based sense annotation and disambiguation

-
**Linguistic Data** to run the server can be download here:

https://github.com/TALP-UPC/FreeLing/tree/master/data

You can download just the ones you need, and put them inside /data like this:

/data/en

/data/pt

