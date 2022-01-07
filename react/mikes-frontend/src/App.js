import Backdrop from './components/Backdrop';
import Todo from './components/Todo';

function App() {
  return (
    <div>

      <h1>My Todos</h1>
      <div>
        <Todo text="Learning React"/>
        <Todo text="Still Learning React"/>
        <Todo text="Done Learning React"/>
      </div>

    </div>
  );
}

export default App;
