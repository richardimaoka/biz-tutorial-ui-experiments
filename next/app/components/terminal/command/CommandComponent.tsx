import styles from "./CommandComponent.module.css";
import { CommandPrompt } from "./CommandPrompt";
import { CommandStringStatic } from "./CommandStringStatic";
import { CommandStringAnimation } from "./CommandStringAnimation";

interface Props {
  command: string;
  animate?: boolean;
}

export function CommandComponent(props: Props) {
  return (
    <div className={styles.component}>
      <pre>
        <CommandPrompt />
        {props.animate ? (
          <CommandStringAnimation command={props.command} />
        ) : (
          <CommandStringStatic command={props.command} />
        )}
      </pre>
    </div>
  );
}
