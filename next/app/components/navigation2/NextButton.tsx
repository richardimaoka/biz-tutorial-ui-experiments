import Link from "next/link";
import { ChevronDownIcon } from "../icons/ChevronDownIcon";
import styles from "./NextButton.module.css";

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
        <div>next</div>
        <ChevronDownIcon />
      </div>
      <div className={`${styles.text} ${styles.desktop}`}>
        <div>Next</div>
        <div>Step</div>
      </div>
    </button>
  </Link>
);
