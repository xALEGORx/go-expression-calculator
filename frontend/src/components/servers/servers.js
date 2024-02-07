import React, { Component } from "react";
import apiService from "../../api";
import Server from "./server";

export default class Servers extends Component {
    constructor(props) {
        super(props);

        this.state = {
            agents: []
        };
    }

    componentDidMount() {
        this.getAgentList()
    }

    getAgentList() {
        apiService.getAgents().then(
            response => {
                this.setState({ agents: response.data.data })
            }
        )
    }

    render() {
        return (
            <div>
                <div className="w-1/2 mb-5">
                    <h1 className="text-black text-xl mb-1">Список серверов</h1>

                    <div className="flex flex-col gap-2">
                        {this.state.agents.map(agent => {
                            return (
                                <Server agent={agent} />
                            )
                        })}
                    </div>
                </div>
            </div>
        );
    }
}
