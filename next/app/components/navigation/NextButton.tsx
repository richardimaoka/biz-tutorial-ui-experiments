import Link from "next/link";
import styles from "./style.module.css";
import { ArrowRightIcon } from "../icons/ArrowRightIcon";

interface NextButtonProps {
  href: string;
}

export const NextButton = ({ href }: NextButtonProps) => (
  <Link
    href={href}
    replace /* TODO: replace parameter not working.... it adds a new entry in Chrome's history*/
  >
    <button className={styles.next}>
      <div className={`${styles.text} ${styles.smartphone}`}>
        <ArrowRightIcon />
      </div>
      <div className={`${styles.text} ${styles.desktop}`}>Next</div>
    </button>
  </Link>
);
