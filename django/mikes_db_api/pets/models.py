from django.db import models

class Species(models.Model):
    name = models.CharField(max_length=50)

class Pet(models.Model):
    name = models.CharField(max_length=25)
    species = models.ForeignKey(Species, on_delete=models.CASCADE)
    
    def __str__(self):
        return f'{self.name} the {self.species.name}'
