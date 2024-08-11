# delus

delus is a Go program that processes a list of domains or subdomains, removes invalid TLDs (Top-Level Domains), and optionally appends a custom string. This special edition is inspired by Lucas, incorporating unique features and flexibility for domain processing.

## Features

- Remove invalid TLDs from domains and subdomains.
- Add a custom string to the end of the cleaned domain, a feature nicknamed "Lucas' Touch".
- Specify the number of words to remove from the end of a domain.
- Forcefully remove parts of a domain even if they seem valid.
- Output the cleaned domains to a specified file or display them in the console.

## Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/LucasOneZ/delus.git
    ```
   
2. **Navigate to the project directory**:
    ```bash
    cd delus
    ```
   
3. **Build the program**:
    ```bash
    go build -o delus main.go
    ```

## Usage

You can run the program with various command-line flags to customize its behavior.

### Basic Usage

```bash
./delus -file=domains.txt
```


Examples

1. Cleaning domains and adding a custom string with Lucas' Touch:

./delus -file=domains.txt -add=<tld or string>

from:

lucas@kali:~/tools/delus$ cat subs
example.comsds
lucas.example.com
example.com.lmao

```bash
output:

lucas@kali:~/tools/delus$ go run main.go -file=subs -add=example.com   

Cleaned: example.com.example.com
Cleaned: example.example.com
Cleaned: lucas.example.com
```

2. Removing the last part of the domain:

./delus -file=domains.txt -removecount=1

```bash
lucas@kali:~/tools/delus$ ./delus -file=domains.txt -removecount=1
Cleaned: example
Cleaned: example.com
Cleaned: lucas.example
```

3. Forcefully removing the last part and adding a custom string:

./delus -file=domains.txt -removecount=1 -force=true -add=<tld or string>

```bash
lucas@kali:~/tools/delus$ ./delus -file=subs -removecount=2 -force=true -add=example.com   
Cleaned: example.example.com
Cleaned: lucas.example.com
```

4. Output cleaned domains to a file:

./delus -file=domains.txt -output=delus.txt

```bash
lucas@kali:~/tools/delus$ ./delus -file=domains.txt -output=delus.txt -verbose                 
Cleaned: example
Cleaned: example.com
Cleaned: lucas.example.com
lucas@kali:~/tools/delus$ cat delus.txt 
example
example.com
lucas.example.com 
```

# Contributing

This is my first tool uploaded to GitHub, and I'm excited to see how it evolves! Contributions are not just welcome but encouraged. If you have ideas for improvements or new features, feel free to fork this repository and submit a pull request. Let's make this tool even better, together with Lucas's guiding inspiration!
