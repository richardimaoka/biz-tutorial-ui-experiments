import { css } from "@emotion/react";
import Link from "next/link";
import styles from "./style.module.css";

interface PrevButtonProps {
  href: string;
}

export const PrevButton = ({ href }: PrevButtonProps) => (
  <Link href={href}>
    <button className={styles.prev}>prev</button>
  </Link>
);
