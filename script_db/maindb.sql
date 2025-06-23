-- WARNING: This schema is for context only and is not meant to be run.
-- Table order and constraints may not be valid for execution.

CREATE TABLE public.article_tags (
  article_id uuid NOT NULL,
  tag_id uuid NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  CONSTRAINT article_tags_pkey PRIMARY KEY (article_id, tag_id),
  CONSTRAINT article_tags_article_id_fkey FOREIGN KEY (article_id) REFERENCES public.articles(id),
  CONSTRAINT article_tags_tag_id_fkey FOREIGN KEY (tag_id) REFERENCES public.tags(id)
);
CREATE TABLE public.article_views (
  article_id uuid NOT NULL,
  user_id uuid,
  ip_address inet,
  user_agent text,
  referrer text,
  country character varying,
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  created_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  CONSTRAINT article_views_pkey PRIMARY KEY (id),
  CONSTRAINT article_views_article_id_fkey FOREIGN KEY (article_id) REFERENCES public.articles(id),
  CONSTRAINT article_views_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.profiles(id)
);
CREATE TABLE public.articles (
  author_id uuid NOT NULL,
  title character varying NOT NULL CHECK (length(title::text) >= 5),
  slug character varying NOT NULL UNIQUE CHECK (slug::text ~ '^[a-z0-9-]+$'::text),
  subtitle character varying,
  content text NOT NULL CHECK (length(content) >= 10),
  excerpt text,
  cover_image_url text,
  cover_image_alt text,
  published_at timestamp with time zone,
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  reading_time integer DEFAULT 1 CHECK (reading_time > 0),
  word_count integer DEFAULT 0,
  published boolean DEFAULT false,
  featured boolean DEFAULT false,
  allow_comments boolean DEFAULT true,
  views_count integer DEFAULT 0,
  likes_count integer DEFAULT 0,
  comments_count integer DEFAULT 0,
  bookmarks_count integer DEFAULT 0,
  created_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  updated_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  CONSTRAINT articles_pkey PRIMARY KEY (id),
  CONSTRAINT articles_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.profiles(id)
);
CREATE TABLE public.bookmarks (
  user_id uuid NOT NULL,
  article_id uuid NOT NULL,
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  created_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  CONSTRAINT bookmarks_pkey PRIMARY KEY (id),
  CONSTRAINT bookmarks_article_id_fkey FOREIGN KEY (article_id) REFERENCES public.articles(id),
  CONSTRAINT bookmarks_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.profiles(id)
);
CREATE TABLE public.comment_likes (
  user_id uuid NOT NULL,
  comment_id uuid NOT NULL,
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  created_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  CONSTRAINT comment_likes_pkey PRIMARY KEY (id),
  CONSTRAINT comment_likes_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.profiles(id),
  CONSTRAINT comment_likes_comment_id_fkey FOREIGN KEY (comment_id) REFERENCES public.comments(id)
);
CREATE TABLE public.comments (
  article_id uuid NOT NULL,
  author_id uuid NOT NULL,
  parent_id uuid,
  content text NOT NULL CHECK (length(TRIM(BOTH FROM content)) >= 1),
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  likes_count integer DEFAULT 0,
  replies_count integer DEFAULT 0,
  is_edited boolean DEFAULT false,
  is_pinned boolean DEFAULT false,
  created_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  updated_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  CONSTRAINT comments_pkey PRIMARY KEY (id),
  CONSTRAINT comments_article_id_fkey FOREIGN KEY (article_id) REFERENCES public.articles(id),
  CONSTRAINT comments_parent_id_fkey FOREIGN KEY (parent_id) REFERENCES public.comments(id),
  CONSTRAINT comments_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.profiles(id)
);
CREATE TABLE public.follows (
  follower_id uuid NOT NULL,
  following_id uuid NOT NULL,
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  created_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  CONSTRAINT follows_pkey PRIMARY KEY (id),
  CONSTRAINT follows_follower_id_fkey FOREIGN KEY (follower_id) REFERENCES public.profiles(id),
  CONSTRAINT follows_following_id_fkey FOREIGN KEY (following_id) REFERENCES public.profiles(id)
);
CREATE TABLE public.likes (
  user_id uuid NOT NULL,
  article_id uuid NOT NULL,
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  created_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  CONSTRAINT likes_pkey PRIMARY KEY (id),
  CONSTRAINT likes_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.profiles(id),
  CONSTRAINT likes_article_id_fkey FOREIGN KEY (article_id) REFERENCES public.articles(id)
);
CREATE TABLE public.notifications (
  recipient_id uuid NOT NULL,
  actor_id uuid,
  type character varying NOT NULL CHECK (type::text = ANY (ARRAY['like'::character varying, 'comment'::character varying, 'follow'::character varying, 'mention'::character varying, 'article_published'::character varying, 'welcome'::character varying]::text[])),
  title character varying NOT NULL,
  message text NOT NULL,
  action_url text,
  read_at timestamp with time zone,
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  created_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  CONSTRAINT notifications_pkey PRIMARY KEY (id),
  CONSTRAINT notifications_actor_id_fkey FOREIGN KEY (actor_id) REFERENCES public.profiles(id),
  CONSTRAINT notifications_recipient_id_fkey FOREIGN KEY (recipient_id) REFERENCES public.profiles(id)
);
CREATE TABLE public.profiles (
  id uuid NOT NULL,
  username character varying NOT NULL UNIQUE,
  full_name character varying,
  bio text,
  avatar_url text,
  website character varying,
  location character varying,
  twitter_username character varying,
  linkedin_username character varying,
  followers_count integer DEFAULT 0,
  following_count integer DEFAULT 0,
  articles_count integer DEFAULT 0,
  total_likes_received integer DEFAULT 0,
  is_verified boolean DEFAULT false,
  is_featured_writer boolean DEFAULT false,
  created_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  updated_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  instagram_username character varying,
  tiktok_username character varying,
  facebook_username character varying,
  CONSTRAINT profiles_pkey PRIMARY KEY (id),
  CONSTRAINT profiles_id_fkey FOREIGN KEY (id) REFERENCES auth.users(id)
);
CREATE TABLE public.tag_follows (
  user_id uuid NOT NULL,
  tag_id uuid NOT NULL,
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  created_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  CONSTRAINT tag_follows_pkey PRIMARY KEY (id),
  CONSTRAINT tag_follows_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.profiles(id),
  CONSTRAINT tag_follows_tag_id_fkey FOREIGN KEY (tag_id) REFERENCES public.tags(id)
);
CREATE TABLE public.tags (
  name character varying NOT NULL UNIQUE CHECK (name::text ~ '^[a-zA-Z0-9\s&+-]+$'::text),
  slug character varying NOT NULL UNIQUE CHECK (slug::text ~ '^[a-z0-9-]+$'::text),
  description text,
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  color character varying DEFAULT '#6366f1'::character varying CHECK (color::text ~ '^#[0-9a-fA-F]{6}$'::text),
  articles_count integer DEFAULT 0,
  followers_count integer DEFAULT 0,
  created_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  updated_at timestamp with time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  CONSTRAINT tags_pkey PRIMARY KEY (id)
);