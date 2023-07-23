#!/usr/bin/env python3

import tkinter as tk
from pathlib import Path
from tkinter import ttk, Menu
from tkinter.messagebox import showinfo, showerror
from tkinterdnd2 import DND_FILES, TkinterDnD
import tkinter.filedialog as fd
import sys

class App(TkinterDnD.Tk):
    def __init__(self):
        super().__init__()

        # configure the root window
        self.title("Rotationg rows in places")
        self.geometry("620x430")

        data_folder = Path("../picture/")
        file_to_open = data_folder / "Avatar_picture.png"

        p1 = tk.PhotoImage(file=file_to_open)

        self.iconphoto(False, p1)

        self.config(menu=MenuBar(self))

        # label

        self.label = ttk.Label(self, text="Drag and drop the file in txt format.")
        self.label.grid(row=0, column=0,  padx=6, pady=6, sticky='ew')

        # Action button
        self.a_button = ttk.Button(self, text="Rotating")
        self.a_button['command'] = self.a_button_clicked
        self.a_button.grid(row=0, column=1,  padx=6, pady=6)


        #Creating a field for dropping files
        self.lbd = tk.Listbox(self, width=100, height=20,)
        self.lbd.insert(1, "Drag files to here")

        # register the listbox as a drop target
        self.lbd.drop_target_register(DND_FILES)
        self.lbd.dnd_bind('<<Drop>>', lambda e: self.lbd.insert(tk.END, e.data))

        self.lbd.grid(row=1, column=0, columnspan=2, sticky='EW', padx=5, pady=5)

        # self.config(menu=MenuBar(self)

        # Delete button
        self.d_button = ttk.Button(self, text="Delete this")
        self.d_button['command'] = self.d_button_clicked
        self.d_button.grid(row=2,column=1, padx=6, pady=6)


    def a_button_clicked(self):
        showinfo(title="Information",
                 message="Hello, Tkinter!")
        self.save_file()

    def d_button_clicked(self):
        selection = self.lbd.curselection()
        self.lbd.delete(selection[0])

    def save_file(self):
        try:
            contents = self.text.get(1.0, tk.END)
            new_file = fd.asksaveasfile(title="Save file", defaultextension=".txt",
                                        filetypes=(("Rotating rows file", "*.txt"),))
            if new_file:
                new_file.write(contents)
                new_file.close()
        except AttributeError:
            showerror(title="Erorr!",
                     message="Hello, Tkinter!")


class MenuBar(tk.Menu):
    def __init__(self, parent):
        tk.Menu.__init__(self, parent)

        filemenu = tk.Menu(self, tearoff=False)
        self.add_cascade(label="File", underline=0, menu=filemenu)
        filemenu.add_command(label="New", command=self.callback)
        filemenu.add_separator()
        filemenu.add_command(label="Exit", underline=1, command=self.quit)

        helpmenu = tk.Menu(self, tearoff=False)
        self.add_cascade(label="Help", menu=helpmenu)
        helpmenu.add_command(label="Help", command=self.help)
        helpmenu.add_command(label="About...", command=self.callback)


    def quit(self):
        sys.exit(0)

    def callback(self):
        print("called the callback!")

    def help(self):
        showinfo(title="Information",
                 message="This application reverses a sequence of lines in a txt or docx extension file. To get the result, drag the desired file into the area, and click the actions button and select the save location.!")


if __name__ == "__main__":
    app = App()
    app.mainloop()
    exit(0)