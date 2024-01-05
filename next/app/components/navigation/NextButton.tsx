"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { ChevronDownIcon } from "../icons/ChevronDownIcon";
import styles from "./NextButton.module.css";

interface Props {
  nextStep: string;
  disabled?: boolean;
}

export function NextButton(props: Props) {
  const pathname = usePathname();

  if (props.disabled) {
    return <></>;
  } else {
    return (
      <Link href={pathname + "?step=" + props.nextStep}>
        <button className={styles.component}>
          <div className={styles.smartphone}>
            <div className={styles.unselectable}>next</div>
            <ChevronDownIcon />
          </div>
          <div className={styles.desktop}>
            <div className={styles.unselectable}>NEXT</div>
            <ChevronDownIcon />
          </div>
        </button>
      </Link>
    );
  }
}
