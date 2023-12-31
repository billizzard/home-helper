### Local file server

This solution is for those who have devices on different platforms (Android, iOS, Windows, Mac) and you need to transfer files between them from time to time (for me this was a headache). By launching the program from the builds folder, you will have a web interface for managing the file. Now all you need to transfer files between devices is a browser. There is no need to install any applications on the devices (only this and only on the main computer). The data is not transferred anywhere, but is stored on your computer (exactly where is described below). There is no interaction with the Internet. But remember, everyone who can access your local network (WiFi, for example) will see these files until you stop the program

### How to use
If you do not plan to develop this platform, but just want to use the solution. Just download the archive from the `builds` folder for your architecture

Then unzip it and go to folder.
You will see two folders and a file

File need to run server. To run it

In linux, open folder in terminal and run
```bash
./file_server
```

In windows, open folder and run
```bash
file_server.exe
```

All files that you upload or download are located in the `user/files`

Folder `src` need to work, don't change this folder

#### Web Interface
After run file_server file you need go to browser to address `http://localhost:9030`. 

If all right you will see
![Alt text](instruction/menu.png?raw=true "Main menu")

Go to `Settings` and you and you will see a local address,
![Alt text](instruction/settings.png?raw=true "Settings")

which you can enter on any device 
connected to the local network to see a folder with shared files, download and upload them (they will all be placed in the `user/files` folder, witch near `file_server` file). 
No need to install any applications. Any device with access to this address will be able 
to see these files

Go to files
![Alt text](instruction/file_item.png?raw=true "File item")

The icon of a file or folder has two areas for clicking; if you click on the icon 
(indicated by a square), you will go to the folder or to the details of the file 
(there is almost empty there for now). If you click on the name (indicated by a circle), 
you will select a file or folder and you will have additional control buttons
(download, rename, delete)


### Develop
To run project in development mode, you need run backend
```bash
make dev-run-back
```

And run react app (located in folder `front`)
```bash
make dev-run-front
```

Go to address
```
http://localhost:3000
```

### Production
To build production builds need to run (linux)
```bash
make prod-build
```

it will create zip archives for different platforms in `builds` folder