import React, { Component } from "react";

export default class Servers extends Component {
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
                <h1>Servers</h1>
            </div>
        );
    }
}
