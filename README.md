# Visual Fanfic

**An app to create content about your favorite work (game, anime, novel).**

Visual Fanfic is with the fans. We support your efforts and devotion for the Games (and the Art in general).

It looks nice to give you some form to obtain more power in expresivity to show your love for the great works to the world.

## What is it?

This app deliver you the tools to create fanfics with visual elements, beautifying them. You can add images and play with styles to get a good theme and produce amazing results.

It uses a web page like form of distribution.

### Config with TOML

More simple and legible config files. You find them in the folder *config*. Here, you see some files: *main.toml* (general config), *characters.toml* (where your characters live) and *dialogue.toml* (script in theatrical production sense).

#### Main

You can config many things (author name, title, about me, contact data) and some of styles. all of needed to customize your theme and present your work. Check *main.toml* for full view of possibilities.

I made it easy for not-developers and standardized it a lot, avoiding many errors to users (an example: I used string type heavily, then all properties require quotation marks). As simple as it's possible.

A quick look:

```toml
# Info
Title = "My Awesome Visual Fanfic"

# Styles
HeaderColor = "#873260"
HeaderFontColor = "#ffffff"
AccentColor = "rebeccapurple" # Web colors: you can use their aliases
DefaultTextColor = "#ffffff"

# Contact data (extensible/deletable properties)
[SocialMedia]
CatFriendsForum = "@your_username"
Phone = "+54911xxxxoooo"
# Continue if you want and it will be added in your footer
```

#### Character

Characters are a representation of one character style. For example, you can have many states of Carla:

```
[[Characters]]
CodeName = "carla_angry" # In your dialogue, use it like reference
VisibleName = "Carla" 
Image = "carla_angry.png"
TextStyle = "n" # Normal style, not bold, not italic
TextColor = "" # Void string applies default color (DefaultTextColor in main.toml)
CharacterBoxColor = ""
DialogBoxColor = ""
NameColor = "#ff0000" # Carla uses color red

[[Characters]]
CodeName = "carla_happy" # Oh, it's different
VisibleName = "Carla"  # Carla keeps her name
Image = "carla_happy.png" # Carla uses a new image
TextStyle = "bi" # Bold and italic
TextColor = ""
CharacterBoxColor = ""
DialogBoxColor = ""
NameColor = "#ff0000" # Carla keeps the color
```

Narrator is a special character (don't have a visible image box, only a dialog box). You can have as many narrators as you like. This way, you can apply the styles you want.

```toml
[[Characters]]
CodeName = "__narrator" # With 2 underlashes '__' in front means "Only text Box"
VisibleName = "" 
Image = ""
TextStyle = "i"
TextColor = ""
CharacterBoxColor = ""
DialogBoxColor = ""
NameColor = ""

[[Characters]]
CodeName = "__witch_singer" # 2 underlashes, it's only a dialog box
VisibleName = "" 
Image = ""
TextStyle = "bi" # See it: Witch Singer uses other style
TextColor = "purple" # Witch Singer has a special color 
CharacterBoxColor = ""
DialogBoxColor = ""
NameColor = ""
```
What is the purpose of it? You could creater a narrator like it:

> My name is Tom Haggard, a musician. This is hard for me to admit, but I'm going to tell you anyway. Here's what happened...

We will explore in next part the dialogue.

#### Dialogue

This step is absurdly easy to understand, but here resides the most poweful tool to build your fanfic. In *dialogue.toml*, you see fragments with two properties: **Cn** (codename) and **Text** (the said text).

```
[[DialogueLines]]
Cn = "__narrator"
Text = "Carla is really frightened. She walks fast when she hears steps behind her."

[[DialogueLines]]
Cn = "carla"
Text = "Who is there?!"
```

As you can see, no need to worry about styles now.

### Images and other resources

It keeps a certain structure in *src* directory (pay attention to it: *output* is similar but it's not where you put your images). A minimal example is delivered with the app.

### Fonts

In *fonts.css* (and some properties in *main.toml*), you can configure the font settings.

### Take the max power

There are two files, *user.css* and *user.js* for advanced users. They can modify all of desired from here (these files are applied in last place).

### Comments to help you

Useful comments in TOML files.

## Execute it

For more compatibility, you can do it (**git** and **go** language are required). 

Clone this repo:

```
git clone git@github.com:EylulBlossomGames/visual-fanfic.git
```

Move to folder, usually:

```
cd visual-fanfic
```

You need **air**:

```
go install github.com/air-verse/air@latest
```

You can run it a single time, or set it up to run automatically on every file change.

Run it once with this command (in root of proyect):

```
go run .
```
Run it in live-reloading mode (on every change) with this command:

```
air
```

## Your feedback

This is WIP  (Work In Progress). Then, please give us your feedback. If you find any bug, report it through the usual channels in Github.

Remember to follow these steps before reporting:

* Gather useful information about your problem.
* Search for keywords related to your problem (has someone already opened an issue about it?).
* Write a descriptive title and explain the situation clearly.
* Be kind to other people.