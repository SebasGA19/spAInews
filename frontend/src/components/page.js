const pageStyle = {
    padding: '2%',
};

export function Page(props) {
    return (
        <div style={pageStyle}>
            {props.children}
        </div>
    );
}