# trnsl8
trnsl8 is a CLI written in Go (using Cobra package) that allows users to translate text and detect the language of input through the use of AWS Translate and AWS Comprehend

<br>


## Installation
```sh
go get -u github.com/gussf/trnsl8
```

<br>
  
## Usage


#### Translation
```sh
# Translates text into a target language
# More information on language codes: https://docs.aws.amazon.com/translate/latest/dg/what-is.html
trnsl8 to <language-code> <text-to-translate>  

# You can input a file to be translated
trnsl8 to <language-code> -f <source-text-file>

# You can save the resulting translation to a file
trnsl8 to <language-code> <text-to-translate> -o <output-file>

# And also both read from and write to files
trnsl8 to <language-code> -f <source-text-file> -o <output-file>
```

#### Detection
```sh
# Detects what language text is in
trnsl8 detect <text-to-detect> 

# You can also input files for detection
trnsl8 detect -f <text-file>
```

<br>

## Examples

```sh
# Translation
trnsl8 to ja "I want this sentence in japanese" 
trnsl8 to en "Ich möchte diesen Satz auf Englisch" -o result.txt
trnsl8 to ro -f ~/source.txt -o ~/result.txt  

# Detection
trnsl8 detect "ce limbă este aceasta?" # hint: its romanian!
trnsl8 detect -f really_big_text.txt
```

<br>

## Next steps
* Use docker?
* Unit tests?
