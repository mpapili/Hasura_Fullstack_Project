from rest_framework import serializers

from pets.models import Pet, Species

class SpeciesSerializer(serializers.ModelSerializer):
    class Meta:
        model = Species
        fields = '__all__'

class PetSerializer(serializers.ModelSerializer):
    species = SpeciesSerializer()

    class Meta:
        model = Pet
        fields = '__all__'