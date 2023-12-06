"use client";

import styles from "./CommandComponent.module.css";
import { CommandPrompt } from "./CommandPrompt";
import { CommandStringStatic } from "./CommandStringStatic";
import { CommandStringAnimation } from "./CommandStringAnimation";
import { useSearchParams } from "next/navigation";

interface Props {
  command: string;
  animate?: boolean;
}

export function CommandComponent(props: Props) {
  const searchParams = useSearchParams();
  const skipAnimation = searchParams.get("skipAnimation") === "true";

  return (
    <div className={styles.component}>
      <pre>
        <CommandPrompt />
        {props.animate && !skipAnimation ? (
          <CommandStringAnimation command={props.command} />
        ) : (
          <CommandStringStatic command={props.command} />
        )}
      </pre>
    </div>
  );
}
