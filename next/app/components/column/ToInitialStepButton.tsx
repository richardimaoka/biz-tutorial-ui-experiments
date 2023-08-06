import Link from "next/link";
import { BackwardFastIcon } from "../icons/BackwardFast";
import styles from "./style.module.css";

export const ToInitialStepButton = () => {
  return (
    <button className={styles.backward}>
      <Link href="/">
        <BackwardFastIcon />
      </Link>
    </button>
  );
};
