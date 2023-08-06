import Link from "next/link";
import styles from "./style.module.css";

interface NextButtonProps {
  href: string;
}

export const NextButton = ({ href }: NextButtonProps) => (
  <Link href={href} className={styles.next}>
    <button>next</button>
  </Link>
);
