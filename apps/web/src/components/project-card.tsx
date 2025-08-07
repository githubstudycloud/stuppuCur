import { Badge } from '@company/ui';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@company/ui';
import { formatRelativeTime } from '@company/utils';
import type { Project } from '@company/types';

interface ProjectCardProps {
  project: Project;
}

const statusColors = {
  planning: 'bg-blue-100 text-blue-800',
  active: 'bg-green-100 text-green-800',
  completed: 'bg-gray-100 text-gray-800',
  archived: 'bg-yellow-100 text-yellow-800',
  cancelled: 'bg-red-100 text-red-800',
};

const priorityColors = {
  low: 'bg-gray-100 text-gray-600',
  medium: 'bg-blue-100 text-blue-600',
  high: 'bg-orange-100 text-orange-600',
  urgent: 'bg-red-100 text-red-600',
};

export function ProjectCard({ project }: ProjectCardProps) {
  return (
    <Card className="hover:shadow-lg transition-shadow">
      <CardHeader>
        <div className="flex items-start justify-between">
          <div>
            <CardTitle className="text-lg">{project.name}</CardTitle>
            <CardDescription className="mt-1">
              {project.description}
            </CardDescription>
          </div>
          <div className="flex flex-col gap-1">
            <span className={`inline-flex items-center rounded-full px-2 py-1 text-xs font-medium ${statusColors[project.status]}`}>
              {project.status}
            </span>
            <span className={`inline-flex items-center rounded-full px-2 py-1 text-xs font-medium ${priorityColors[project.priority]}`}>
              {project.priority}
            </span>
          </div>
        </div>
      </CardHeader>
      <CardContent>
        <div className="space-y-3">
          {/* Tags */}
          <div className="flex flex-wrap gap-1">
            {project.tags.map((tag) => (
              <span
                key={tag}
                className="inline-flex items-center rounded-md bg-gray-50 px-2 py-1 text-xs font-medium text-gray-600 ring-1 ring-inset ring-gray-500/10"
              >
                {tag}
              </span>
            ))}
          </div>

          {/* Dates */}
          <div className="text-sm text-muted-foreground space-y-1">
            {project.startDate && (
              <div>
                <span className="font-medium">Start:</span> {formatRelativeTime(project.startDate)}
              </div>
            )}
            {project.endDate && (
              <div>
                <span className="font-medium">End:</span> {formatRelativeTime(project.endDate)}
              </div>
            )}
            <div>
              <span className="font-medium">Updated:</span> {formatRelativeTime(project.updatedAt)}
            </div>
          </div>

          {/* Team info */}
          <div className="flex items-center justify-between text-sm text-muted-foreground">
            <span>{project.teamMembers.length} team members</span>
            <span>Owner: {project.ownerId}</span>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}