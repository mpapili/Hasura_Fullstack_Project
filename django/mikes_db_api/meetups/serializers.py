from rest_framework import serializers

from meetups.models import Meetup

class MeetupSerializer(serializers.ModelSerializer):
    class Meta:
        model = Meetup
        fields = '__all__'
