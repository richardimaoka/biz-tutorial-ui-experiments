import Link from "next/link";
import styles from "./style.module.css";
import { ArrowRightIcon } from "../icons/ArrowRightIcon";

interface NextButtonProps {
  href: string;
}

export const NextButton = ({ href }: NextButtonProps) => (
  <Link href={href}>
    <button className={styles.next}>
      <div className={`${styles.text} ${styles.smartphone}`}>
        <ArrowRightIcon />
      </div>
      <div className={`${styles.text} ${styles.desktop}`}>Next</div>
    </button>
  </Link>
);
