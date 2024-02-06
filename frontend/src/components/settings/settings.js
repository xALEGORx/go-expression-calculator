import React, { Component } from "react";

export default class Settings extends Component {
    constructor(props) {
        super(props);

        this.state = {
            content: ""
        };
    }

    componentDidMount() {
        this.setState({ content: [] })
    }

    render() {
        return (
            <div>
                <h1>Settings</h1>
            </div>
        );
    }
}
