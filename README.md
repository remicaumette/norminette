# norminette

Because my norminette doesn't work.

## install

```bash
mkdir ~/.bin
curl -L https://github.com/remicaumette/norminette/releases/download/v1.3.0/norminette -o ~/.bin/norminette
chmod a+x ~/.bin/norminette
echo "export PATH=\"\$PATH:~/.bin\"" >> ~/.zshrc
echo "alias norminette=~/.bin/norminette" >> ~/.zshrc
source ~/.zshrc
```

## get started

- help:
```
norminette -help
```

- usage:
```
norminette file.c ./src
```
