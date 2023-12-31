"use client";

import styles from "./CommandComponent.module.css";
import { CommandPrompt } from "./CommandPrompt";
import { CommandWaitingExecution } from "./CommandWaitingExecution";
import { CommandTypingAnimation } from "./CommandTypingAnimation";
import { useSearchParams } from "next/navigation";
import { CommandAlreadyExecuted } from "./CommandAlreadyExecuted";

interface Props {
  command: string;
  animate?: boolean;
}

export function CommandComponent(props: Props) {
  const searchParams = useSearchParams();
  const skipAnimation = searchParams.get("skipAnimation") === "true";

  if (props.animate) {
    // If this command has `props.animate = true`, then the command is suppsedly the last command waiting for execution
    if (skipAnimation) {
      // This might feel weird, but `props.animate = true` && `skipAnimation (search params) = true`
      // will only animate the flickering trail, indicating that the command is waiting for execution
      return (
        <div className={styles.component}>
          <pre>
            <CommandPrompt />
            <CommandWaitingExecution command={props.command} />
          </pre>
        </div>
      );
    } else {
      // `props.animate = true` && `!skipAnimation (search params)`
      // Full animation = typing + flickering trail
      return (
        <div className={styles.component}>
          <pre>
            <CommandPrompt />
            <CommandTypingAnimation command={props.command} />
          </pre>
        </div>
      );
    }
  } else {
    // If `!props.animate` supposedly the command is already executed so no animation at all
    <div className={styles.component}>
      <pre>
        <CommandPrompt />
        <CommandAlreadyExecuted command={props.command} />
      </pre>
    </div>;
  }
}
