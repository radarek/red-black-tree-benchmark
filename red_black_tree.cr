class RedBlackTree(T)
  class Node(T)
    property :color
    property :key
    property :left
    property :right
    property :parent

    enum Color
      Red
      Black
    end

    def initialize(key : T, color : Color)
      @left = @right = @parent = NilNode.instance
      @color = color
      @key = key
    end

    def black?
      @color.black?
    end

    def red?
      @color.red?
    end
  end

  class NilNode
    property :left, :right, :parent, :color, :key

    def initialize
      @parent = self
    end

    def self.instance
      @@instance ||= new
    end

    def color
      Node::Color::Black
    end

    def black?
      true
    end

    def red?
      false
    end
    
    def left
      self
    end

    def right
      self
    end

    def key
      raise "I don't exist"
    end

    def nil?
      true
    end
  end

  property :root
  property :size

  def initialize
    @root = NilNode.instance
    @size = 0
  end

  def add(key)
    insert(Node(T).new(key, Node::Color::Red))
  end

  def minimum(x = root)
    while !x.left.nil?
      x = x.left
    end
    return x
  end

  def maximum(x = root)
    while !x.right.nil?
      x = x.right
    end
    return x
  end

  def successor(x)
    if !x.right.nil?
      return minimum(x.right)
    end
    y = x.parent
    while !y.nil? && x == y.right
      x = y
      y = y.parent
    end
    return y
  end

  def predecessor(x)
    if !x.left.nil?
      return maximum(x.left)
    end
    y = x.parent
    while !y.nil? && x == y.left
      x = y
      y = y.parent
    end
    return y
  end

  def inorder_walk(x = root)
    x = self.minimum
    while !x.nil?
      yield x.key
      x = successor(x)
    end
  end

  def reverse_inorder_walk(x = root)
    x = self.maximum
    while !x.nil?
      yield x.key
      x = predecessor(x)
    end
  end

  def search(key, x = root)
    while !x.nil? && x.key != key
      x = (key < x.key ? x.left : x.right)
    end
    return x
  end

  def empty?
    return self.root.nil?
  end

  def black_height(x = root)
    height = 0
    while !x.nil?
      x = x.left
      height +=1 if x.nil? || x.black?
    end
    return height
  end

  def delete(z)
    y = (z.left.nil? || z.right.nil?) ? z : successor(z)
    x = y.left.nil? ? y.right : y.left
    x.parent = y.parent

    if y.parent.nil?
      self.root = x
    else
      if y == y.parent.left
        y.parent.left = x
      else
        y.parent.right = x
      end
    end

    z.key = y.key if y != z

    if y.black?
      delete_fixup(x)
    end

    self.size -= 1
    return y
  end

  def insert(x)
    insert_helper(x)

    x.color = Node::Color::Red
    while x != root && x.parent.red?
      if x.parent == x.parent.parent.left
        y = x.parent.parent.right
        if !y.nil? && y.red?
          x.parent.color = Node::Color::Black
          y.color = Node::Color::Black
          x.parent.parent.color = Node::Color::Red
          x = x.parent.parent
        else
          if x == x.parent.right
            x = x.parent
            left_rotate(x)
          end
          x.parent.color = Node::Color::Black
          x.parent.parent.color = Node::Color::Red
          right_rotate(x.parent.parent)
        end
      else
        y = x.parent.parent.left
        if !y.nil? && y.red?
          x.parent.color = Node::Color::Black
          y.color = Node::Color::Black
          x.parent.parent.color = Node::Color::Red
          x = x.parent.parent
        else
          if x == x.parent.left
            x = x.parent
            right_rotate(x)
          end
          x.parent.color = Node::Color::Black
          x.parent.parent.color = Node::Color::Red
          left_rotate(x.parent.parent)
        end
      end
    end
    root.color = Node::Color::Black
  end

  private def insert_helper(z)
    y = NilNode.instance
    x = root
    while !x.nil?
      y = x
      x = z.key < x.key ? x.left : x.right
    end
    z.parent = y
    if y.nil?
      self.root = z
    else
      z.key < y.key ? y.left = z : y.right = z
    end
    self.size += 1
  end

  private def left_rotate(x)
    raise "x.right is nil!" if x.right.nil?
    y = x.right
    x.right = y.left
    y.left.parent = x if !y.left.nil?
    y.parent = x.parent
    if x.parent.nil?
      self.root = y
    else
      if x == x.parent.left
        x.parent.left = y
      else
        x.parent.right = y
      end
    end
    y.left = x
    x.parent = y
  end

  private def right_rotate(x)
    raise "x.left is nil!" if x.left.nil?
    y = x.left
    x.left = y.right
    y.right.parent = x if !y.right.nil?
    y.parent = x.parent
    if x.parent.nil?
      self.root = y
    else
      if x == x.parent.left
        x.parent.left = y
      else
        x.parent.right = y
      end
    end
    y.right = x
    x.parent = y
  end

  private def delete_fixup(x)
    while x != root && x.black?
      if x == x.parent.left
        w = x.parent.right
        if w.red?
          w.color = Node::Color::Black
          x.parent.color = Node::Color::Red
          left_rotate(x.parent)
          w = x.parent.right
        end
        if w.left.black? && w.right.black?
          w.color = Node::Color::Red
          x = x.parent
        else
          if w.right.black?
            w.left.color = Node::Color::Black
            w.color = Node::Color::Red
            right_rotate(w)
            w = x.parent.right
          end
          w.color = x.parent.color
          x.parent.color = Node::Color::Black
          w.right.color = Node::Color::Black
          left_rotate(x.parent)
          x = root
        end
      else
        w = x.parent.left
        if w.red?
          w.color = Node::Color::Black
          x.parent.color = Node::Color::Red
          right_rotate(x.parent)
          w = x.parent.left
        end
        if w.right.black? && w.left.black?
          w.color = Node::Color::Red
          x = x.parent
        else
          if w.left.black?
            w.right.color = Node::Color::Black
            w.color = Node::Color::Red
            left_rotate(w)
            w = x.parent.left
          end
          w.color = x.parent.color
          x.parent.color = Node::Color::Black
          w.left.color = Node::Color::Black
          right_rotate(x.parent)
          x = root
        end
      end
    end
    x.color = Node::Color::Black
  end
end

