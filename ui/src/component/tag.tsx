import { LucideIcon } from "lucide-react";

type TagProps = {
  icon: LucideIcon;
  value: string;
};
const Tag = (props: TagProps) => {
  return (
    <div className="flex items-center gap-1">
      <props.icon className="text-secondary size-4" />
      <span className="text-secondary text-sm">{props.value}</span>
    </div>
  );
};

export default Tag;
