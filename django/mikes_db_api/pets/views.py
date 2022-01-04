from rest_framework.decorators import api_view
from rest_framework.response import Response
from pets.models import Pet
from pets.serializers import PetSerializer


@api_view(['GET'])
def pet_detail(request, pk):
    
    print('debug, running pet detail!')
    pet = Pet.objects.get(id=pk)
    print(f'found, {pet}!')
    serializer = PetSerializer(pet, many=False)
    print(f'debug {serializer.data}')
    return Response({'message': serializer.data})