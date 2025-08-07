import { Button } from '@company/ui';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@company/ui';
import { formatDate, formatCurrency } from '@company/utils';
import { APP_NAME, APP_DESCRIPTION } from '@company/config';
import { ProjectCard } from '../components/project-card';

const mockProjects = [
  {
    id: '1',
    name: 'E-commerce Platform',
    description: 'Modern shopping experience with React and Node.js',
    status: 'active' as const,
    priority: 'high' as const,
    startDate: '2024-01-15T00:00:00Z',
    endDate: '2024-06-15T00:00:00Z',
    createdAt: '2024-01-01T00:00:00Z',
    updatedAt: '2024-02-01T00:00:00Z',
    ownerId: 'user1',
    teamMembers: [],
    tags: ['react', 'nodejs', 'ecommerce'],
  },
  {
    id: '2',
    name: 'Analytics Dashboard',
    description: 'Real-time data visualization and reporting',
    status: 'planning' as const,
    priority: 'medium' as const,
    startDate: '2024-03-01T00:00:00Z',
    createdAt: '2024-02-15T00:00:00Z',
    updatedAt: '2024-02-15T00:00:00Z',
    ownerId: 'user2',
    teamMembers: [],
    tags: ['dashboard', 'analytics', 'charts'],
  },
];

export default function HomePage() {
  return (
    <div className="container mx-auto px-4 py-12">
      {/* Hero Section */}
      <section className="text-center mb-16">
        <h1 className="text-4xl font-bold tracking-tight mb-4">
          Welcome to {APP_NAME}
        </h1>
        <p className="text-xl text-muted-foreground mb-8 max-w-2xl mx-auto">
          {APP_DESCRIPTION}
        </p>
        <div className="flex gap-4 justify-center">
          <Button size="lg">Get Started</Button>
          <Button variant="outline" size="lg">
            Learn More
          </Button>
        </div>
      </section>

      {/* Features Section */}
      <section className="mb-16">
        <h2 className="text-3xl font-bold text-center mb-12">Key Features</h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <Card>
            <CardHeader>
              <CardTitle>Shared Components</CardTitle>
              <CardDescription>
                Reusable UI components across all applications
              </CardDescription>
            </CardHeader>
            <CardContent>
              <p className="text-sm text-muted-foreground">
                Consistent design system with TypeScript support and modern styling.
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Type Safety</CardTitle>
              <CardDescription>
                End-to-end TypeScript with shared type definitions
              </CardDescription>
            </CardHeader>
            <CardContent>
              <p className="text-sm text-muted-foreground">
                Catch errors early with comprehensive type checking across packages.
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Developer Experience</CardTitle>
              <CardDescription>
                Modern tooling with hot reload and optimized builds
              </CardDescription>
            </CardHeader>
            <CardContent>
              <p className="text-sm text-muted-foreground">
                Turbo for fast builds, ESLint, Prettier, and comprehensive testing.
              </p>
            </CardContent>
          </Card>
        </div>
      </section>

      {/* Stats Section */}
      <section className="mb-16">
        <Card>
          <CardHeader>
            <CardTitle>Platform Statistics</CardTitle>
            <CardDescription>
              Real-time metrics from our monorepo
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
              <div className="text-center">
                <div className="text-2xl font-bold">4</div>
                <div className="text-sm text-muted-foreground">Packages</div>
              </div>
              <div className="text-center">
                <div className="text-2xl font-bold">3</div>
                <div className="text-sm text-muted-foreground">Applications</div>
              </div>
              <div className="text-center">
                <div className="text-2xl font-bold">{formatCurrency(99999)}</div>
                <div className="text-sm text-muted-foreground">Development Value</div>
              </div>
              <div className="text-center">
                <div className="text-2xl font-bold">{formatDate(new Date())}</div>
                <div className="text-sm text-muted-foreground">Last Updated</div>
              </div>
            </div>
          </CardContent>
        </Card>
      </section>

      {/* Projects Section */}
      <section>
        <h2 className="text-3xl font-bold text-center mb-12">Active Projects</h2>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {mockProjects.map((project) => (
            <ProjectCard key={project.id} project={project} />
          ))}
        </div>
      </section>
    </div>
  );
}