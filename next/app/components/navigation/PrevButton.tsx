"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { ChevronUpIcon } from "../icons/ChevronUpIcon";
import styles from "./PrevButton.module.css";

interface Props {
  prevStep: string;
  disabled?: boolean;
}

export function PrevButton(props: Props) {
  const pathname = usePathname();

  if (props.disabled) {
    return <></>;
  } else {
    return (
      <Link href={pathname + "?step=" + props.prevStep + "&skipAnimation=true"}>
        <button className={styles.component}>
          <div className={styles.smartphone}>
            <div>prev</div>
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
}
