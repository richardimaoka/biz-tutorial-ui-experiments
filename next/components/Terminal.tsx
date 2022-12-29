import { css } from "@emotion/react";
import React from "react";
import { useEffect, useRef } from "react";
import { useState } from "react";
import { Command } from "./Command";
import { Output } from "./Output";

export interface Command {
  command: string;
  output: string;
}

interface State {
  state: "command writing" | "wait command execution" | "output writing";
  stepAt: number;
  commandWrittenLength: number;
}

interface TerminalProps {
  list: Command[];
}

export const Terminal = ({ list }: TerminalProps): JSX.Element => {
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
          setState({ ...state, state: "wait command execution" });
        }
        break;
      case "wait command execution":
        // console.log("wait command execution");
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
            <Command
              command={command}
              writtenLength={state.commandWrittenLength}
            />
          </div>
        );
      case "wait command execution":
        return (
          <div ref={ref}>
            <Command command={command} writtenLength={command.length} />
          </div>
        );
      case "output writing":
        const output = list[state.stepAt].output;
        return (
          <>
            <Command command={command} writtenLength={command.length} />
            <div ref={ref}>
              <Output output={output} />
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
          background-color: #3a3a3a;
          height: 300px;
          overflow-y: auto;
        `}
      >
        {list
          //render up to previous step's element
          .filter((_, index) => index < state.stepAt)
          .map((elem, index) => (
            <React.Fragment key={index}>
              <Command
                command={elem.command}
                writtenLength={elem.command.length}
              />
              <Output output={elem.output} />
            </React.Fragment>
          ))}
        <LastElement />
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
