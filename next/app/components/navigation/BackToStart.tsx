import Link from "next/link";

import styles from "./style.module.css";

export const BackToStart = () => (
  <Link href="/" className={styles.back}>
    <div className={styles.back}>back to start</div>
  </Link>
);
