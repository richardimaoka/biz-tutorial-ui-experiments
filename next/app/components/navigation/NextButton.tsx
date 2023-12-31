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
  props.disabled = false;

  if (props.disabled) {
    return <></>;
  } else {
    return (
      <Link href={pathname + "?step=" + props.nextStep}>
        <button className={styles.component}>
          <div className={styles.smartphone}>
            <div>next</div>
            <ChevronDownIcon />
          </div>
          <div className={styles.desktop}>
            <div>NEXT</div>
            <ChevronDownIcon />
          </div>
        </button>
      </Link>
    );
  }
}
