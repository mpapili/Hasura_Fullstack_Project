from rest_framework.decorators import api_view
from rest_framework import generics
from rest_framework.response import Response

from meetups.models import Meetup
from meetups.serializers import MeetupSerializer

@api_view(['POST'])
def add_meetup(request):
    '''
    create a new meetup record from POST data
    '''
    print(request.data)
    # this is not true REST but let's pretend it is
    new_meetup = Meetup.objects.create(
        title=request.data['title'],
        image=request.data['image'],
        address=request.data['address'],
        description=request.data['description'],
    )
    return Response({'message': f'Meetup Created with id {new_meetup.id}'})


class MeetupView(generics.ListAPIView):

    queryset = Meetup.objects.all()
    serializer_class = MeetupSerializer
