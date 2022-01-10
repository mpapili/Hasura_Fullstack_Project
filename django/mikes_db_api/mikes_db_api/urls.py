"""mikes_db_api URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/4.0/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path
from django.views.generic import TemplateView

from meetups import views as meetup_views
from pets import views as pet_views

urlpatterns = [
    path('admin/', admin.site.urls),
    
    # pets
    path('pets/<int:pk>/', pet_views.pet_detail, name='pet_detail'),

    # meetups
    path('meetups/add', meetup_views.add_meetup, name='add_meetup'),
    path('meetups', meetup_views.MeetupView.as_view()),

]
