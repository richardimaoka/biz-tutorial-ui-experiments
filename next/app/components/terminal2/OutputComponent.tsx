import styles from "./OutputComponent.module.css";

interface Props {
  output: string;
}

export function OutputComponent(props: Props) {
  return (
    <div className={styles.component}>
      <pre>
        <code>{props.output}</code>
      </pre>
    </div>
  );
}
