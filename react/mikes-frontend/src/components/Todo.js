function Todo(props) {

    function deleteHandler() {
        window.alert("You clicked button with title " + props.text);
    }

    return (
        <div className='card'>
            <h2>{props.text}</h2>
            <button className='btn' onClick={deleteHandler}>Delete</button>
        </div>
    )   
}

export default Todo;