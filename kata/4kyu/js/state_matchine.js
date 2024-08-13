let states = new Map();

states.set("CLOSE_WAIT",
    [{
        input: "APP_CLOSE",
        next: "LAST_ACK"
    }])

states.set("LAST_ACK",
    [{
        input: "RCV_ACK",
        next: "CLOSED"
    }])

states.set("TIME_WAIT",
    [{
        input: "APP_TIMEOUT",
        next: "CLOSED"
    }])

states.set("FIN_WAIT_2",
    [{
        input: "RCV_FIN",
        next: "TIME_WAIT"
    }])

states.set("CLOSING",
    [{
        input: "RCV_ACK",
        next: "TIME_WAIT"
    }])

states.set("FIN_WAIT_1",
    [{
        input: "RCV_FIN",
        next: "CLOSING"
    }, {
        input: "RCV_FIN_ACK",
        next: "TIME_WAIT"
    }, {
        input: "RCV_ACK",
        next: "FIN_WAIT_2"
    }])

states.set("ESTABLISHED",
    [{
        input: "APP_CLOSE",
        next: "FIN_WAIT_1"
    }, {
        input: "RCV_FIN",
        next: "CLOSE_WAIT"
    }])

states.set("SYN_SENT",
    [{
        input: "RCV_SYN",
        next: "SYN_RCVD"
    }, {
        input: "RCV_SYN_ACK",
        next: "ESTABLISHED"
    }, {
        input: "APP_CLOSE",
        next: "CLOSED"
    }])

states.set("SYN_RCVD",
    [{
        input: "APP_CLOSE",
        next: "FIN_WAIT_1"
    }, {
        input: "RCV_ACK",
        next: "ESTABLISHED"
    }])

states.set("CLOSED",
    [{
        input: "APP_PASSIVE_OPEN",
        next: "LISTEN"
    }, {
        input: "APP_ACTIVE_OPEN",
        next: "SYN_SENT"
    }]);

states.set("LISTEN",
    [{
        input: "RCV_SYN",
        next: "SYN_RCVD"
    }, {
        input: "APP_SEND",
        next: "SYN_SENT"
    }, {
        input: "APP_CLOSE",
        next: "CLOSED"
    }])

function traverseTCPStates(eventList) {
    let state = "CLOSED";  // initial state, always

    for (let i = 0; i < eventList.length; i++) {
        let nx = states.get(state);

        let find = false;
        for (let j = 0; j < nx.length; j++) {
            if (nx[j].input === eventList[i]) {
                state = nx[j].next;
                find = true;
            }
        }
        if (!find) {
            return "ERROR";
        }
    }

    return state;
}

console.log(traverseTCPStates(["APP_ACTIVE_OPEN","RCV_SYN_ACK","RCV_FIN"]));
