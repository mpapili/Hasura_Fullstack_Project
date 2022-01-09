import classes from './Card.module.css';

function Card(props) {
    // we want to inject the JSX we wrap in here
    // the wrapped content will get passed in as a special component called "props.children!"
    return <div className={classes.card}>
        {props.children}
    </div>
}

export default Card;