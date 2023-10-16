import Link from "next/link";
import { ChevronDownIcon } from "../icons/ChevronDownIcon";
import styles from "./PrevButton.module.css";

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
        <div>prev</div>
        <ChevronDownIcon />
      </div>
      <div className={`${styles.text} ${styles.desktop}`}>
        <div>Prev</div>
        <div>Step</div>
      </div>
    </button>
  </Link>
);
