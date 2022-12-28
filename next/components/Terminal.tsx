import { css } from "@emotion/react";
import { useEffect, useRef } from "react";
import { useState } from "react";

const expected = [
  `docker build -t cmd1 -f Dockerfile.cmd1 .
docker run --rm cmd1`,
  "abc",
  `docker build -t cmd2 -f Dockerfile.cmd2 .
docker run --rm cmd2`,
  "abc def",
  `docker build -t cmd3 -f Dockerfile.cmd3 .
docker run --rm cmd3`,
  "abc",
  `docker build -t cmd4 -f Dockerfile.cmd4 .
docker run --rm cmd4`,
  "abc def",
  `docker build -t cmd5 -f Dockerfile.cmd5 .
docker run --rm cmd5`,
  `/home/your_username`,
  `docker build -t cmd6 -f Dockerfile.cmd6 .
docker run --rm cmd6`,
  `/home/your_username`,
  `docker pull nginx
docker inspect nginx`,
  `"Config": {
    "Cmd": [
        "nginx",
        "-g",
        "daemon off;"
    ]
}`,
  `docker run nginx:1.23.1`,
  `2022/09/03 07:14:58 [notice] 1#1: using the "epoll" event method
2022/09/03 07:14:58 [notice] 1#1: nginx/1.23.1
2022/09/03 07:14:58 [notice] 1#1: built by gcc 10.2.1 20210110 (Debian 10.2.1-6)
2022/09/03 07:14:58 [notice] 1#1: OS: Linux 5.10.102.1-microsoft-standard-WSL2
2022/09/03 07:14:58 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576
2022/09/03 07:14:58 [notice] 1#1: start worker processes
2022/09/03 07:14:58 [notice] 1#1: start worker process 31
2022/09/03 07:14:58 [notice] 1#1: start worker process 32
2022/09/03 07:14:58 [notice] 1#1: start worker process 33
2022/09/03 07:14:58 [notice] 1#1: start worker process 34
2022/09/03 07:14:58 [notice] 1#1: start worker process 35
2022/09/03 07:14:58 [notice] 1#1: start worker process 36
2022/09/03 07:14:58 [notice] 1#1: start worker process 37
2022/09/03 07:14:58 [notice] 1#1: start worker process 38`,
  `docker build -t cmd-nginx -f Dockerfile.cmd-nginx .
docker run --rm cmd-nginx`,
  `2022/09/03 05:48:46 [notice] 7#7: start worker process 13
2022/09/03 05:48:46 [notice] 7#7: start worker process 14
2022/09/03 05:48:46 [notice] 7#7: start worker process 15`,
  `^C`,
];

interface Command {
  command: string;
  output?: string;
}

const expected2 = [
  {
    command: `docker build -t cmd1 -f Dockerfile.cmd1 .
docker run --rm cmd1`,
    output: "abc",
  },
  {
    command: `docker build -t cmd2 -f Dockerfile.cmd2 .
docker run --rm cmd2`,
    output: "abc def",
  },
  {
    command: `docker build -t cmd3 -f Dockerfile.cmd3 .
docker run --rm cmd3`,
    output: "abc",
  },
  {
    command: `docker build -t cmd4 -f Dockerfile.cmd4 .
docker run --rm cmd4`,
    output: "abc def",
  },
  {
    command: `docker build -t cmd5 -f Dockerfile.cmd5 .
docker run --rm cmd5`,
    output: `/home/your_username`,
  },
  {
    command: `docker build -t cmd6 -f Dockerfile.cmd6 .
docker run --rm cmd6`,
    output: `/home/your_username`,
  },
  {
    command: `docker pull nginx
docker inspect nginx`,
    output: `"Config": {
    "Cmd": [
        "nginx",
        "-g",
        "daemon off;"
    ]
}`,
  },
  {
    command: `docker run nginx:1.23.1`,
    output: `2022/09/03 07:14:58 [notice] 1#1: using the "epoll" event method
2022/09/03 07:14:58 [notice] 1#1: nginx/1.23.1
2022/09/03 07:14:58 [notice] 1#1: built by gcc 10.2.1 20210110 (Debian 10.2.1-6)
2022/09/03 07:14:58 [notice] 1#1: OS: Linux 5.10.102.1-microsoft-standard-WSL2
2022/09/03 07:14:58 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576
2022/09/03 07:14:58 [notice] 1#1: start worker processes
2022/09/03 07:14:58 [notice] 1#1: start worker process 31
2022/09/03 07:14:58 [notice] 1#1: start worker process 32
2022/09/03 07:14:58 [notice] 1#1: start worker process 33
2022/09/03 07:14:58 [notice] 1#1: start worker process 34
2022/09/03 07:14:58 [notice] 1#1: start worker process 35
2022/09/03 07:14:58 [notice] 1#1: start worker process 36
2022/09/03 07:14:58 [notice] 1#1: start worker process 37
2022/09/03 07:14:58 [notice] 1#1: start worker process 38`,
  },
  {
    command: `docker build -t cmd-nginx -f Dockerfile.cmd-nginx .
docker run --rm cmd-nginx`,
    output: `2022/09/03 05:48:46 [notice] 7#7: start worker process 13
2022/09/03 05:48:46 [notice] 7#7: start worker process 14
2022/09/03 05:48:46 [notice] 7#7: start worker process 15`,
  },
];

interface State {
  state: "command writing" | "wait command execution" | "output writing";
  stepAt: number;
  commandWrittenLength: number;
}

export const Terminal = (): JSX.Element => {
  const [state, setState] = useState<State>({
    state: "command writing",
    stepAt: 0,
    commandWrittenLength: 0,
  });
  const ref = useRef(null);

  const cssRunnable = css`
    background-color: green;
  `;
  const cssNotRunnable = css`
    background-color: gray;
  `;

  const cssTerminalCommand = css`
    margin: 1px 0px;
    padding: 4px;
    background-color: #3a3a3a;
    color: white;
    border-bottom: 1px solid white;
  `;
  const cssTerminalOutput = css`
    margin: 1px 0px;
    padding: 4px;
    background-color: #3a3a3a;
    color: white;
    border-bottom: 1px solid white;
  `;

  // const updateCommandCharAt = () => {
  //   const command = expected2[state.stepAt].command;
  //   const durationMills = 10000;
  //   const numCycles = 4;
  //   const charsPerCycle = command.length / numCycles;
  //   let nextCharAt = Math.min(
  //     Math.round(state.commandCharAt + charsPerCycle),
  //     Math.round(command.length - 1)
  //   );
  //   console.log(
  //     `updateCommandCharAt state.commandCharAt=${state.commandCharAt} command.length=${command.length} nextCharAt=${nextCharAt}`
  //   );
  //   if (state.commandCharAt < command.length - 1) {
  //     setState({ ...state, commandCharAt: nextCharAt });
  //     setTimeout(updateCommandCharAt, durationMills / numCycles);
  //   } else {
  //     setState({ ...state, state: "wait command execution" });
  //   }
  // };

  useEffect(() => {
    if (ref.current) {
      ref.current.scrollIntoView();
    }
    switch (state.state) {
      case "command writing":
        console.log("command writing");
        const command = expected2[state.stepAt].command;
        const durationMills = 3000;
        const numCycles = 10;
        const charsPerCycle = command.length / numCycles;

        let nextWrittenLength = Math.min(
          Math.round(state.commandWrittenLength + charsPerCycle),
          command.length
        );
        console.log(
          `updateCommandCharAt state.commandWrittenLength=${state.commandWrittenLength} command.length=${command.length} nextWrittenLength=${nextWrittenLength}`
        );

        if (state.commandWrittenLength < command.length) {
          setTimeout(() => {
            setState({ ...state, commandWrittenLength: nextWrittenLength });
          }, durationMills / numCycles);
        } else {
          setState({ ...state, state: "wait command execution" });
        }
        break;
      case "wait command execution":
        console.log("wait command execution");
        break;
      case "output writing":
        console.log("output writing");
        if (state.stepAt < expected2.length - 1) {
          setTimeout(() => {
            setState({
              ...state,
              stepAt: state.stepAt + 1,
              state: "command writing",
            });
          }, 300);
        }
        break;
      default:
        const _exhaustiveCheck: never = state.state;
        return _exhaustiveCheck;
    }
  });

  const onClick = () => {
    setState({ ...state, state: "output writing" });
  };

  return (
    <div>
      <div
        css={css`
          background-color: #3a3a3a;
          height: 300px;
          overflow-y: auto;
        `}
      >
        {expected2
          //render up to previous step's element
          .filter((_, index) => index < state.stepAt)
          .map((elem, index) => (
            <>
              <pre css={cssTerminalCommand} key={"command" + index}>
                <code>{elem.command}</code>
              </pre>
              <pre css={cssTerminalOutput} key={"output" + index}>
                <code>{elem.output}</code>
              </pre>
            </>
          ))}
        {state.state === "command writing" ? (
          <pre
            css={cssTerminalCommand}
            key={"command" + state.stepAt}
            ref={ref}
          >
            <code>
              {expected2[state.stepAt].command.substring(
                0,
                state.commandWrittenLength
              )}
            </code>
          </pre>
        ) : state.state === "wait command execution" ? (
          <>
            <pre
              css={cssTerminalCommand}
              key={"command" + state.stepAt}
              ref={ref}
            >
              <code>{expected2[state.stepAt].command}</code>
            </pre>
          </>
        ) : state.state === "output writing" ? (
          <>
            <pre css={cssTerminalCommand} key={"command" + state.stepAt}>
              <code>{expected2[state.stepAt].command}</code>
            </pre>
            <pre
              css={cssTerminalOutput}
              key={"output" + state.stepAt}
              ref={ref}
            >
              <code>{expected2[state.stepAt].output}</code>
            </pre>
          </>
        ) : (
          <></>
        )}
      </div>
      <button
        disabled={state.state !== "wait command execution"}
        type="button"
        onClick={onClick}
      >
        run
      </button>
    </div>
  );
};
