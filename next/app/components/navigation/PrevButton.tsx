"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { ChevronUpIcon } from "../icons/ChevronUpIcon";
import styles from "./PrevButton.module.css";
interface Props {
  prevStep: string;
}

export function PrevButton(props: Props) {
  const pathname = usePathname();

  return (
    <Link href={pathname + "?step=" + props.prevStep}>
      <button className={styles.component}>
        <div className={styles.smartphone}>
          <div>next</div>
          <ChevronUpIcon />
        </div>
        <div className={styles.desktop}>
          <div>PREV</div>
          <ChevronUpIcon />
        </div>
      </button>
    </Link>
  );
}
