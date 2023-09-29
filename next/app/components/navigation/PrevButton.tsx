import Link from "next/link";
import styles from "./style.module.css";
import { ArrowLeftIcon } from "../icons/ArrowLeftIcon";

interface PrevButtonProps {
  href: string;
}

export const PrevButton = ({ href }: PrevButtonProps) => (
  <Link
    href={href}
    replace /* TODO: replace parameter not working.... it adds a new entry in Chrome's history*/
  >
    <button className={styles.prev}>
      <div className={`${styles.text} ${styles.smartphone}`}>
        <ArrowLeftIcon />
      </div>
      <div className={`${styles.text} ${styles.desktop}`}>Prev</div>
    </button>
  </Link>
);
