from django.db import models

# Create your models here.
class Meetup(models.Model):
    title = models.CharField(max_length=60)
    address = models.CharField(max_length=60)
    image = models.CharField(max_length=255)
    description = models.TextField(max_length=255)
