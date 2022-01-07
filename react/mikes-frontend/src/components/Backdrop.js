function Backdrop(props) {
    // we're using a PROP to pass a function!!!
    return <div className='backdrop' onClick={props.onClick} />
}

export default Backdrop;