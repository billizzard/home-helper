### Develop
Need to install sass compiler, for it need download release

`https://github.com/sass/dart-sass/releases`

Unpack archive, copy content to ~/.local/bin

And then add folder to PATH

For it add to ~/.bash_profile

`export PATH="$PATH:/home/mi/.local/bin"`

to check success installation

`sass --version`

Then just run

`sass public/assets/scss/main.scss public/assets/css/main.css`


### Start app

`npm start`