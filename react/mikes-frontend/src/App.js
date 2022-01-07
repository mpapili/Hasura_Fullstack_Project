import { Route, Routes } from 'react-router-dom';

import Todo from './components/Todo';
import AllMeetupsPage from './pages/AllMeetups';
import FavoritesPage from './pages/Favorites';
import NewMeetupPage from './pages/NewMeetup';
import MainNavigation from './components/layout/MainNavigation';

function App() {

  // localhost:3000 is our main

  return (
    <div>
      <MainNavigation />
      <Routes>
        <Route path='/' element={<AllMeetupsPage />} />
        <Route path='/new-meetup' element={<NewMeetupPage />} />
        <Route path='/favorites' element={<FavoritesPage />} />
      </Routes>


      {/*
      DEMO COMPONENTS:
      <h1>My Todos</h1>
      <div>
        <Todo text="Learning React"/>
        <Todo text="Still Learning React"/>
        <Todo text="Done Learning React"/>
      </div>
      */}

    </div>
  );
}

export default App;
