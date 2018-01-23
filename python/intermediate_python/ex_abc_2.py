from abc import ABCMeta, abstractmethod


class Vehicle(object):
    __metaclass__ = ABCMeta

    @abstractmethod
    def change_gear(self):
        pass

    @abstractmethod
    def start_engine(self):
        pass


class Car(Vehicle):
    def __init__(self, make, model, color):
        self.make = make
        self.model = model
        self.color = color

    def change_gear(self):
        print("Changing Gear")

    def start_engine(self):
        print("Starting engine")


car = Car("Toyota", "Corolla", "black")
print car.make
