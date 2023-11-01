import styles from "./CommandComponent.module.css";
import { CommandPrompt } from "./CommandPrompt";
import { CommandStringStatic } from "./CommandStringStatic";
import { CommandStringAnimation } from "./CommandStringAnimation";

interface Props {
  command: string;
  animate?: boolean;
  completedCallback?: () => void;
}

export function CommandComponent(props: Props) {
  return (
    <div className={styles.component}>
      <pre>
        <CommandPrompt />
        {props.animate ? (
          <CommandStringAnimation
            command={props.command}
            completedCallback={props.completedCallback}
          />
        ) : (
          <CommandStringStatic command={props.command} />
        )}
      </pre>
    </div>
  );
}
