import { useRouter } from "next/navigation";
import styles from "./TerminalTabs.module.css";
import { TerminalTab } from "./TerminalTab";

interface Props {
  tabs: {
    name: string;
    href: string; //TODO, calculate href??
  }[];
  selectTab: string;
}

export function TerminalTabs(props: Props) {
  return (
    <div className={styles.component}>
      {props.tabs.map((x) => (
        <TerminalTab
          key={x.name}
          href={x.href}
          isSelected={x.name === props.selectTab}
          name={x.name}
        />
      ))}
    </div>
  );
}
