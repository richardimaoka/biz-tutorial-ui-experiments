import styles from "./style.module.css";

export const StepDisplay = ({ step }: { step: string }) => (
  <div className={styles.step}>step = {step}</div>
);
