import React from 'react'

export default function Server({ agent }) {
    const statusServer = {
        "connected": "-[#08c708]",
        "disconnected": "-[#c70808]",
    };

    const statusTextServer = "text" + statusServer[agent.status]
    const statusBgServer = "bg" + statusServer[agent.status]

    return (
        <div className="flex gap-2">
            <span className={`inline-block w-12 h-12 rounded-full ${statusBgServer}`}></span>
            <div className="flex flex-col justify-between">
                <span className="text-black text-xl font-semibold">{agent.agent_id} <span className="text-gray text-sm font-normal">{agent.last_ping}</span></span>
                <span className={`${statusTextServer}`}>{agent.status}</span>
            </div>
        </div>
    )
}
