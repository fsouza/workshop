class Widget(object):

    def __init__(self, x, y):
        self.x = x
        self.y = y

    def dimensions(self):
        return self.x, self.y


class Label(Widget):

    def __init__(self, text, x, y):
        super(Label, self).__init__(x, y)
        self.text = text

if __name__ == "__main__":
    l = Label("")
    l.dimensions()
