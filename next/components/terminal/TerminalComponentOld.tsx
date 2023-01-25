import { css } from "@emotion/react";
import React, { useEffect, useRef, useState } from "react";
import { CommandComponent } from "./CommandComponent";
import { CommandOutputComponent } from "./CommandOutputComponent";

export interface Command {
  command: string;
  output: string;
}

interface State {
  state: "command writing" | "command ready" | "output writing";
  stepAt: number;
  commandWrittenLength: number;
}

interface TerminalComponentOldProps {
  list: Command[];
}

export const TerminalComponentOld = ({
  list,
}: TerminalComponentOldProps): JSX.Element => {
  const [state, setState] = useState<State>({
    state: "command writing",
    stepAt: 0,
    commandWrittenLength: 0,
  });
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (ref.current) {
      ref.current.scrollIntoView();
    }
    switch (state.state) {
      case "command writing":
        // console.log("command writing");
        const command = list[state.stepAt].command;
        if (state.commandWrittenLength < command.length) {
          setTimeout(() => {
            setState({
              ...state,
              commandWrittenLength: state.commandWrittenLength + 6,
            });
          }, 30);
        } else {
          setState({ ...state, state: "command ready" });
        }
        break;
      case "command ready":
        // console.log("command ready");
        break;
      case "output writing":
        // console.log("output writing");
        if (state.stepAt < list.length - 1) {
          setTimeout(() => {
            setState({
              stepAt: state.stepAt + 1,
              state: "command writing",
              commandWrittenLength: 0,
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

  const LastElement = (): JSX.Element => {
    const command = list[state.stepAt].command;
    switch (state.state) {
      case "command writing":
        return (
          <div ref={ref}>
            <CommandComponent
              command={command}
              writtenLength={state.commandWrittenLength}
            />
          </div>
        );
      case "command ready":
        return (
          <div ref={ref}>
            <CommandComponent
              command={command}
              writtenLength={command.length}
            />
          </div>
        );
      case "output writing":
        const output = list[state.stepAt].output;
        return (
          <>
            <CommandComponent
              command={command}
              writtenLength={command.length}
            />
            <div ref={ref}>
              <CommandOutputComponent output={output} />
            </div>
          </>
        );
      default:
        const _exhaustiveCheck: never = state.state;
        return _exhaustiveCheck;
    }
  };

  return (
    <div>
      <div
        css={css`
          background-color: #1e1e1e;
          height: 300px;
          overflow-y: auto;
        `}
      >
        {list
          //render up to previous step's element
          .filter((_, index) => index < state.stepAt)
          .map((elem, index) => (
            <React.Fragment key={index}>
              <CommandComponent
                command={elem.command}
                writtenLength={elem.command.length}
              />
              <CommandOutputComponent output={elem.output} />
            </React.Fragment>
          ))}
        <LastElement />
      </div>
      <div
        css={css`
          background-color: #1e1e1e;
        `}
      >
        <button
          disabled={state.state !== "command ready"}
          type="button"
          onClick={onClick}
        >
          run
        </button>
      </div>
    </div>
  );
};
